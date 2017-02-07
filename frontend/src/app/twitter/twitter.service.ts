import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Rx';
import { Tweet } from './tweet';

@Injectable()
export class TwitterService {

  private apiUrl: string = 'http://localhost:8090/api/v1/twitter/timeline';

  constructor(private http: Http) { }

  getTimeline(nbEvents: number){
    return this.http.get([this.apiUrl, nbEvents].join('/'))
      .map(res => <Tweet[]>res.json());
  }
}
