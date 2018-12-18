import { Injectable } from '@angular/core';
import { Cat } from './cat';
import { CATS } from './mock-cats';
import { Observable, of } from "rxjs";
import { MessageService } from "./message.service";

@Injectable({
  providedIn: 'root'
})
export class CatService {

  constructor(private messageService: MessageService) { }

  getCats(): Observable<Cat[]> {
    // TODO: send the message _after_ fetching the cats
    this.messageService.add('CatService: fetched cats');
    return of(CATS);
  }

}
