import { ApiService } from './../api/api.service';
import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Rx';

@Injectable()
export class YunService {

  constructor(private http: Http,private apiService : ApiService) {
  }

  getRoomStat(room){
      return this.http.get(this.apiService.getUrl(room))
        .map(res => res.json());
  }

  

}
