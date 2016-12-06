import { SortPipe } from './../pipes/sort.pipe';
import { LunchplaceService } from './lunchplace.service';
import { LunchplaceComponent } from './lunchplace.component';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

@NgModule({
    declarations: [ LunchplaceComponent ],
    exports: [ LunchplaceComponent ],
    providers: [ LunchplaceService],
    imports: [ CommonModule ]
})
export class LunchplaceModule {

}