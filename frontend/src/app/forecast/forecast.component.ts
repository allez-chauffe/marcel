import { Component, OnInit,ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'forecast',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './forecast.component.html',
  styleUrls: ['./forecast.component.scss']
})
export class ForecastComponent implements OnInit {

  public numberslist : number[];

  constructor() { 
    this.numberslist = [0,1,2];
  }

  ngOnInit() {
  }

}
