import { SortPipe } from './../pipes/sort.pipe';
import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Rx';

@Injectable()
export class GithubService {

  api: string = 'https://api.github.com/users/';
  params: string = '/events';
  users: any[] = [];
  usersContribusion: any[] = [];

  constructor(private http: Http, private sort: SortPipe) {
    this.users.push(
      'Gillespie59',
      'GwennaelBuchet',
      'T3kstiil3',
      'RemiEven',
      'looztra',
      'a-cordier',
      'wadendo',
      'NathanDM',
      'Antoinephi',
      'cluster',
      'yyekhlef',
      'gdrouet',
      'Kize'
    );
  }

  getTopContributors(){
    return Observable
      .from(this.users)
      .flatMap((users) => this.getUserContributions(users))
      .map((usersContri) => {
        let nbContributions = 0;

        usersContri
          .filter(contrib => contrib.type === 'PushEvent' && contrib.created_at.indexOf(this.getDate()) !== -1)
          .map(contri => nbContributions += contri.payload.commits.length);

        this.usersContribusion.push({
          user : usersContri[0].actor.display_login,
          contributions : nbContributions
        });
        return this.sort.transform(this.usersContribusion, 'contributions');
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

  private getDate(){
    let today = new Date();
    return `${today.getFullYear()}-${("0" + (today.getMonth() + 1)).slice(-2)}`;
  }

}
