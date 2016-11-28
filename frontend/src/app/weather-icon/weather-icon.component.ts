import {Component, OnInit} from "@angular/core";
import {ViewEncapsulation} from '@angular/core';

import {WeatherService} from '../weather/weather.service';
import {FirstletterupperPipe} from '../pipes/firstletterupper.pipe';

@Component({
  selector: 'weather-icon',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './weather-icon.component.html',
  styleUrls: ['./weather-icon.component.scss']
})

export class WeatherIconComponent implements OnInit{

  public current_weather_type : String = "";
  public weather_types : Array<Object> = [];
  public weather_object : any;
  public weather_temp : number = 0;
  public number_list : number = 1; 

  constructor(public weatherService: WeatherService){
    this.weather_types.push("sun-shower");
    this.weather_types.push("thunder-storm");
    this.weather_types.push("cloudy");
    this.weather_types.push("sunny");
    this.weather_types.push("rainy");
    this.weather_types.push("flurries");
  }

  ngOnInit(){
    console.log('Init weather icon');
    this.getWeather();
    setInterval(() => {
      this.getWeather();
    }, 180000);
  }

  getWeather(){
    this.weatherService.getCurrent()
    .subscribe(res => {
      this.weather_temp = Math.round(res.list[this.number_list].main.temp);
      this.setWeatherType(res.list[this.number_list].weather[0].main);
    });
  }

  randomType(){
    var randomInt = Math.floor(Math.random() * this.weather_types.length);
    return this.weather_types[randomInt];
  }

  setWeatherType(type) {
    if(!type){
      type = this.randomType();
    }
    this.current_weather_type = type;
  }

}
