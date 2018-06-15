import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { InsultsComponent } from './insults.component';

describe('InsultsComponent', () => {
  let component: InsultsComponent;
  let fixture: ComponentFixture<InsultsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ InsultsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(InsultsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
