import { Injectable } from '@angular/core';
import { Http } from '@angular/http';

@Injectable()
export class CalendarService {

  public api_url : String = "http://10.0.10.63:8090/api/v1/agenda/incoming/5";

  constructor(public http : Http) {}

  getEvents(){
    return this.http
      .get(this.api_url+'?json_callback=JSON_CALLBACK')
      .map(function(res){
        const response = res.json();
        return response.items;
      });
  }

}
