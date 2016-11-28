import {Injectable} from '@angular/core';
import {Http} from '@angular/http';
import {ApiService} from './../api/api.service';

@Injectable()
export class WeatherService {

  public api : string;

  constructor(public http:Http,public apiService:ApiService) {
    this.api = this.apiService.getApi('weather').url;
  }

  getCurrent() {
    console.log(this.api);
      return this.http.get(this.api)
        .map(res => res.json());
  }

  getForecast() {
      return this.http.get(this.api)
        .map(res => res.json());
  }
}
