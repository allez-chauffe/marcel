import { Injectable } from '@angular/core';
import { Http } from '@angular/http';

@Injectable()
export class CalendarService {

  public api_url : String = "app/calendar/agenda.json";

  constructor(public http : Http) {
  }

  getEvents(){
    return this.http
      .get(this.api_url+'?json_callback=JSON_CALLBACK')
      .map(function(res){
        const response = res.json();
        return response.items;
      });
  }

}
