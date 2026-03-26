extern crate taubyte_sdk;
use taubyte_sdk::event::Event;

#[no_mangle]
pub fn tencent(event: Event) -> u32 {
    let http = event.http().unwrap();
    _ = http.write("hello tencent".as_bytes()).unwrap();
    0
}
