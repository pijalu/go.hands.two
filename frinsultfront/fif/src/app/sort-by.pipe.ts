import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'sortBy'
})
export class SortByPipe implements PipeTransform {

  transform(ar: Array<any>, args?: any): Array<any> {
    if (ar !== undefined) {
      ar.sort((a: any, b: any) => {
        if (a[args] < b[args]) {
          return 1;
        } else if (a[args] > b[args]) {
          return -1;
        } else {
          return 0;
        }
      });
    }
    return ar;
  }

}
