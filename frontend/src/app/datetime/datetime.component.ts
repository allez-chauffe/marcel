import { DateTime } from './datetime.model';
import {Component, OnInit} from "@angular/core";
import {ViewEncapsulation} from '@angular/core';

@Component({
  selector: 'datetime',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './datetime.component.html',
  styleUrls: ['./datetime.component.scss']
})
export class DateTimeComponent implements OnInit{

  public date: DateTime;

  constructor() { }

  ngOnInit(){
    this.loopDate();
    setInterval(() => {
      this.loopDate();
    }, 10000);
  }

  loopDate(){
    this.date = new DateTime(new Date());
  }
}
