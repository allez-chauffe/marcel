import {Injectable} from '@angular/core';
import {Http} from '@angular/http';
import {ApiService} from './../api/api.service';

@Injectable()
export class WeatherService {

  public apiForecast : String = "http://api.openweathermap.org/data/2.5/forecast";
  public apiCurrent : String = "http://api.openweathermap.org/data/2.5/weather";
  public key : String = "macle";
  public key_url : String;
  public requestUrl : String;
  public params : String = "&q=Lille,Fr&units=metric";

  constructor(public http:Http,public apiService:ApiService) {
    this.key = this.apiService.getKey('weather');
    this.key_url = '?APPID='+this.key;
  }

  getCurrent() {
      return this.http.get(this.apiCurrent+''+this.key_url+''+this.params)
        .map(res => res.json());
  }

  getForecast() {
      return this.http.get(this.apiForecast+''+this.key_url+''+this.params)
        .map(res => res.json());
  }
}
