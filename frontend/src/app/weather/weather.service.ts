import {Injectable} from '@angular/core';
import {Http} from '@angular/http';
import {ApiService} from './../api/api.service';
import { environment } from '../../environments/environment';

@Injectable()
export class WeatherService {

  public api : string;
  public key : string;
  public apiForecast : String = "http://api.openweathermap.org/data/2.5/forecast";

  constructor(public http:Http,public apiService:ApiService) {
    if (environment.production && 0) {
      this.api = this.apiService.getApi('weather').url;
    }else{
      this.key = this.apiService.getApi('weather').key;
      this.api = this.apiForecast + '?APPID=' + this.key + '&q=Lille,Fr&units=metric';
    }
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
