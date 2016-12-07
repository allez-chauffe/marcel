import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { environment } from '../../environments/environment';

@Injectable()
export class CalendarService {

  public api_url : String;

  constructor(public http : Http) {
    this.api_url = environment.agendaUrl;
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
