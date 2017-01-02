import { ApiService } from './../api/api.service';
import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Rx';

@Injectable()
export class SoundtouchService {

  now_playing : string = ':8090/now_playing';

  constructor(private http: Http,private apiService : ApiService) {}

  getNowPlaying(){
    console.log('on passe la !!!');
    console.log(this.apiService.getUrl('soundtouch') + this.now_playing);
      return this.http.get(this.apiService.getUrl('soundtouch') + this.now_playing)
        .map(res => res.json())
        .map(res => {
          console.log('la soundtouch !');
          console.log(res);
          return res;
        });
  }

}
