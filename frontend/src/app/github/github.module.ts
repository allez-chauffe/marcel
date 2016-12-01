import { SortPipe } from './../pipes/sort.pipe';
import { GithubService } from './github.service';
import { GithubComponent } from './github.component';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

@NgModule({
    declarations: [ GithubComponent ],
    exports: [ GithubComponent ],
    providers: [ GithubService, SortPipe],
    imports: [ CommonModule ]
})
export class GithubModule {

}