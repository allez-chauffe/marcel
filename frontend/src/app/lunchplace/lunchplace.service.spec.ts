/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { LunchplaceService } from './lunchplace.service';

describe('Service: Lunchplace', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [LunchplaceService]
    });
  });

  it('should ...', inject([LunchplaceService], (service: LunchplaceService) => {
    expect(service).toBeTruthy();
  }));
});
