import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Insult } from './insult';
import { environment } from '../environments/environment.prod';



@Injectable({
  providedIn: 'root'
})
export class FrinsultService {
  private url = environment.insultUrl;
  
  constructor(private http: HttpClient) { }
  
  getInsults(): Observable<Insult[]> {
    return this.http.get<Insult[]>(this.url)
  }

  vote(url: string, callback) {
    this.http.post(url, null).subscribe(
    res => {
      callback()
    },
    err => {
      console.log("Error occured", err);
    })
  }

  upvote(id: number, callback) {
    this.vote(this.url+"/upvote/"+id, callback)
  }

  downvote(id: number, callback) {
    this.vote(this.url+"/downvote/"+id, callback)
  }
}
