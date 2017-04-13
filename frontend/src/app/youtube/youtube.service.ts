import { Http, Response } from '@angular/http';
import { Injectable } from '@angular/core';
import { ApiService } from '../api/api.service';
import { Observable } from 'rxjs';
import { Subject } from 'rxjs/Subject';

@Injectable()
export class YoutubeService {

  private subject = new Subject<any>();

  constructor(private http: Http, private apiService: ApiService) {

  }

  startSearch(query: String) {
    this.subject.next({query: query});
  }

  getSearch(): Observable<any> {
    return this.subject.asObservable();
  }

  search(query: String) {
    return this.http.get(this.apiService.getApi('youtube').url +'?q=' + query + '&part=snippet&key=' + this.apiService.getApi('youtube').key)
      .map((res: Response) => res.json())
      .map(json => json.items);
  }

}
