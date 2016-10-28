/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { CalendarService } from './calendar.service';

describe('Service: Calendar', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [CalendarService]
    });
  });

  it('should ...', inject([CalendarService], (service: CalendarService) => {
    expect(service).toBeTruthy();
  }));
});
