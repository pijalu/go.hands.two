import { Component, OnInit } from '@angular/core';
import { Insult } from '../insult';
import { FrinsultService } from '../frinsult.service';

@Component({
  selector: 'app-insults',
  templateUrl: './insults.component.html',
  styleUrls: ['./insults.component.css']
})
export class InsultsComponent implements OnInit {
  insults : Insult[]

  constructor(private service: FrinsultService) { }

  ngOnInit() {
    this.getInsults()
  }

  getInsults() {
    this.service.getInsults().subscribe(insults => this.insults = insults)
  }

  upvote(id: number) {
    console.log("upvote", id);  
    this.service.upvote(id, () => this.getInsults())
  }

  downvote(id: number) {    
    console.log("downvote", id);
    this.service.downvote(id, () => this.getInsults())
  }
}
