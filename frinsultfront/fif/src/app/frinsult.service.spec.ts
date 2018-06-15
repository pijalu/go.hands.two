import { TestBed, inject } from '@angular/core/testing';

import { FrinsultService } from './frinsult.service';

describe('FrinsultService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [FrinsultService]
    });
  });

  it('should be created', inject([FrinsultService], (service: FrinsultService) => {
    expect(service).toBeTruthy();
  }));
});
