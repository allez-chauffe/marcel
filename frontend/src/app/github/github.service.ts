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

  getTopContributors() {
    return Observable
      .from(this.users)
      .flatMap(user => this.getUserInfo(user))
      .flatMap(user => this.getUserContributions(user))
      .map(contribsAndUser => {
        let nbContributions = 0;
        let contribs = contribsAndUser.contribs;

        contribs
          .filter(contrib => contrib.type === 'PushEvent' && contrib.created_at.indexOf(this.getDate()) !== -1)
          .map(contri => nbContributions += contri.payload.commits.length);

        this.usersContribusion.push({
          user: contribsAndUser.user.name,
          contributions: nbContributions,
          avatar: contribs[0].actor.avatar_url
        });
        return this.sort.transform(this.usersContribusion, 'contributions');
      });
  }

  getUserInfo(user: string) {
    return this.http.get(`${this.api}${user}`)
      .map(res => {
        let response = res.json();
        return {
          login: user,
          name: response['name']
        }
      });
  }

  getUserContributions(user) {
    return this.http
      .get(this.api + user.login + this.params)
      .map(res => {
        return {
          contribs: res.json(),
          user
        };
      });
  }

  private getDate() {
    let today = new Date();
    return `${today.getFullYear()}-${("0" + (today.getMonth() + 1)).slice(-2)}`;
  }

}
