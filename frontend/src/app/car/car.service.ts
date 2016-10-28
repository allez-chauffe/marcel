import {Injectable} from '@angular/core';
import { Http } from '@angular/http';
import {ApiService} from './../api/api.service';

@Injectable()
export class CarService {

  public api_url : String = "";

  constructor(public http : Http,public apiService:ApiService) {
    this.api_url = apiService.getUrl('xee');
  }

  getCarData(data){
    return this.http
      .get(this.api_url+data)
      .map(res => res.json());
  }
}
