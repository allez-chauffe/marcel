/* tslint:disable:no-unused-variable */

import { TestBed, async, inject } from '@angular/core/testing';
import { GithubService } from './github.service';

describe('Service: Github', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [GithubService]
    });
  });

  it('should ...', inject([GithubService], (service: GithubService) => {
    expect(service).toBeTruthy();
  }));
});
