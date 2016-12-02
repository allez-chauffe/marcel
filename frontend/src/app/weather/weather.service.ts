import {Injectable} from '@angular/core';
import {Http} from '@angular/http';
import { environment } from '../../environments/environment';

@Injectable()
export class WeatherService {

  public api: string;
  public key: string;
  public apiForecast: String;

  constructor(public http: Http) {
    this.key = environment.weatherKey;
    this.api = environment.weatherUrl + '?APPID=' + this.key + '&q=Lille,Fr&units=metric';
  }

  getCurrent() {
      return this.http.get(this.api)
        .map(res => res.json());
  }

  getForecast() {
      return this.http.get(this.api)
        .map(res => res.json());
  }
}
