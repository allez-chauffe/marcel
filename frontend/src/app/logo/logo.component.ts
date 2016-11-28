import { Component, OnInit, ViewEncapsulation } from '@angular/core';

@Component({
  selector: 'logo',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './logo.component.html',
  styleUrls: ['./logo.component.scss']
})
export class LogoComponent implements OnInit {

  constructor() { }

  ngOnInit() {
  }

}
