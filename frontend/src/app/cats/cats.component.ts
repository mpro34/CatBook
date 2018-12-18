import { Component, OnInit } from '@angular/core';
import { Cat } from '../cat';
import { CATS } from '../mock-cats';
import { CatService } from "../cat.service";

@Component({
  selector: 'app-cats',
  templateUrl: './cats.component.html',
  styleUrls: ['./cats.component.css']
})
export class CatsComponent implements OnInit {

  cats: Cat[];
  selectedCat: Cat;

  constructor(private catService: CatService) {}


  ngOnInit() {
    this.getCats();
  }

  getCats(): void {
    this.catService.getCats()
      .subscribe(cats => this.cats = cats);
  }

  onSelect(cat: Cat): void {
    this.selectedCat = cat;
  }

}
