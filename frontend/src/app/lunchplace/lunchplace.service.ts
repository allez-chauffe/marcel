import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Rx';

@Injectable()
export class LunchplaceService {

  api : string = 'http://10.0.10.3:8080/';

  constructor(private http: Http) {

  }

  get_teams_by_orga(orga_id){
    let url = this.api+'organizations/'+orga_id;
    return this.http.get(url)
        .map(res => res.json())
        .map(orga => {
          let teams = [];
          if(orga[0] && orga[0].teams)
            teams = orga[0].teams;
          return teams;
        })
        .flatMap(teams => 
            Observable
            .from(teams)
            .flatMap((team) => this.get_daily_restaurant(team))
        );
  }

  get_daily_restaurant(team){
    let url = this.api+'teams/'+team.id+'/dailyRestaurant/';
    return this.http.get(url)
        .map(res => res.json());
  }

}
