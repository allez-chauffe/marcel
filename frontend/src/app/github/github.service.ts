import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Rx';

@Injectable()
export class GithubService {

  api : string = 'https://api.github.com/users/'
  params : string = '/events'
  users : any[] = []
  usersContribusion : any[] = []

  constructor(public http : Http) {
    this.users.push(
      'T3kstiil3',
      'Gillespie59',
      'GwennaelBuchet',
      'RemiEven',
      'looztra',
      'a-cordier'
    );
  }

  getTopContributors(){
    return Observable
      .from(this.users)
      .flatMap((users) => this.getUserContributions(users))
      .map((usersContri) => {
        let nbContributions = 0;
        
        usersContri.map((contri)=>{
          if(contri.created_at.indexOf('2016-11') !== -1)
            nbContributions++;
        });

        this.usersContribusion.push({
          user : usersContri[0].actor.display_login,
          contributions : nbContributions
        });

        return this.usersContribusion;
      });
  }

  getUserContributions(user){
    return this.http
      .get(this.api + user + this.params)
      .map(function(res){
        let response = res.json();
        return response;
      });
  }



}
