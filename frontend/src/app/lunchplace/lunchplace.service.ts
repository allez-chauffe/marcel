import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Rx';

@Injectable()
export class LunchplaceService {

  api : string = 'http://lunchplace-rct.wip.devlab722.net/api/v1/';

  constructor(private http: Http) {}

  get_teams_by_orga(orga_id){
    let url = this.api+'organizations/'+orga_id;
    return this.http.get(url)
      .map(res => res.json())
      .map(orga => {
        let teams = [];
        
        if(orga[0] && orga[0].teams)
          teams = orga[0].teams;
        
        return {
          orga_id: orga_id,
          teams: teams
        }
      })
  }

  get_orga_daily_restaurant(orga){
    return Observable
      .from(orga.teams)
      .flatMap(team => this.get_team_daily_restaurant(team))
      .toArray()
      .map(teams => {
        return {
            orga: orga,
            teams : teams
          };
      })
  }

  get_team_daily_restaurant(team){
    let url = this.api+'teams/'+team.id+'/dailyRestaurant/';
    return this.http.get(url)
      .map(res => res.json())
      .map(restaurant => {
        team.daily = restaurant;
        return team;
      })
  }

  get_teams_daily(orga_id){
    return Observable
      .from([orga_id])
      .flatMap(orga_id => this.get_teams_by_orga(orga_id))
      .flatMap(orga => this.get_orga_daily_restaurant(orga));
  }
}
