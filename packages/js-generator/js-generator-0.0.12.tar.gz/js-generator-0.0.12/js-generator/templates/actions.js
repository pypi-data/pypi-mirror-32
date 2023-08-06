import HttpService from '../../services/http.service';
import ApiService from '../../services/api.service';

export const SYNC_EXAMPLE = 'EXAMPLE';
export const ASYNC_EXAMPLE = 'ASYNC_EXAMPLE';

/**
 * Async actions
 * */
export const AsyncExampleAction = createAsyncAction(
  ASYNC_EXAMPLE, (name, age) => {
    const options = ApiService.getOptions('exampleEndPoint');
    return HttpService.fetch({...options, body: JSON.stringify({name, age})})
  }
);

/**
* Sync actions
* */
export function SyncExampleAction(payload) {
 return {
   type: SYNC_EXAMPLE,
   payload
 }
}
