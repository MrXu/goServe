import {bootstrap}    from 'angular2/platform/browser';
import 'rxjs/Rx';
import {HTTP_PROVIDERS} from 'angular2/http';
import {TopNavLayout} from './components/TopNavLayout';

bootstrap(TopNavLayout, [HTTP_PROVIDERS]);