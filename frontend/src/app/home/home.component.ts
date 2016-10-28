import {Component, OnInit,ElementRef} from "@angular/core";
import {ViewEncapsulation} from '@angular/core';
import {NgStyle} from '@angular/common';

@Component({
  selector: 'home',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})

export class HomeComponent implements OnInit{

  public _sRotate: string = "1";

  constructor() {
  }

  ngOnInit(){
    console.log('Init Home');
  }
}
