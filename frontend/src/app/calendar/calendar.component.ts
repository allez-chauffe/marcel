import { Component, OnInit } from '@angular/core';
import { ViewEncapsulation } from '@angular/core';
import { CalendarService } from './calendar.service';
import { Observable } from 'rxjs/Rx';

@Component({
  selector: 'calendar',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.scss']
})
export class CalendarComponent implements OnInit {

  public events: any[] = [];

  private timer: number = 1000 * 60 * 60;

  constructor(public calendarService: CalendarService) {

  }

  ngOnInit() {

    setInterval(() => {
      this.calendarService.getEvents()
        .subscribe(o => {
          this.events = o;
        });
    }, this.timer);
  }

}
