fn main () {
    let mut num = 15;

        match num {
            10 => println!("Hello world!"),
            15 => {
                for i in 1..26{
                    println!("Number {}", i)
                }
            }
            _ => println!("")
        }


    /* for i in 0..101 {
        if i % 10 == 0{
            println!("{}",i)
        }
    }*/


    /*while num >= 0 {

        println!("{}",num);
        num -= 1;

    }*/

    /*loop {
        if num == 100 {
            break;
        }

        println!("{}", num);
        num += 1;
    }*/

}
