import { Component, OnInit } from '@angular/core';
import { ViewEncapsulation} from '@angular/core';
import { CalendarService } from './calendar.service';
import { Observable } from 'rxjs/Rx';

@Component({
  selector: 'calendar',
  encapsulation: ViewEncapsulation.None,
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.scss']
})
export class CalendarComponent implements OnInit {

  public events : any[] = [];

  constructor(public calendarService : CalendarService) {

  }

  ngOnInit() {
    console.log('Init Calandar');
    this.calendarService.getEvents()
      .subscribe((o) => {
        this.events = o;
        console.log(this.events);
      });
  }

}
