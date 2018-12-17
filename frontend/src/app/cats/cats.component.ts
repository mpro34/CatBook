import { Component, OnInit } from '@angular/core';
import { Cat } from '../cat';
import { CATS } from '../mock-cats';

@Component({
  selector: 'app-cats',
  templateUrl: './cats.component.html',
  styleUrls: ['./cats.component.css']
})
export class CatsComponent implements OnInit {

  cats = CATS;
  selectedCat: Cat;

  constructor() { }

  ngOnInit() {
  }

  onSelect(cat: Cat): void {
    this.selectedCat = cat;
  }

}
