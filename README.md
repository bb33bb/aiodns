# aiodns
[![](https://img.shields.io/badge/Telegram-Group-blue)](https://t.me/aioCloud)
[![](https://img.shields.io/badge/Telegram-Channel-green)](https://t.me/aioCloud_channel) 

```bash
wget -O aiodns.conf https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/master/accelerated-domains.china.conf
sed -i "s/server=\///" aiodns.conf
sed -i "s/\/114.114.114.114/./" aiodns.conf
```

```csharp
namespace AIODNSTester
{
    public class AIODNSTester
    {
        [DllImport("aiodns.bin", CallingConvention = CallingConvention.Cdecl)]
        public static extern bool aiodns_dial(byte[] chinacon, byte[] chinadns, byte[] otherdns);

        [DllImport("aiodns.bin", CallingConvention = CallingConvention.Cdecl)]
        public static extern bool aiodns_init();

        [DllImport("aiodns.bin", CallingConvention = CallingConvention.Cdecl)]
        public static extern void aiodns_free();

        public static void Main(string[] args)
        {
            Console.WriteLine("[AIODNSTester][aiodns_dial]");
            if (!aiodns_dial(Encoding.UTF8.GetBytes("D:\\china.conf"), Encoding.UTF8.GetBytes("119.29.29.29:53"), Encoding.UTF8.GetBytes("1.1.1.1:53")))
            {
                Console.ReadLine();
                return;
            }
            Console.ReadLine();

            Console.WriteLine("[AIODNSTester][aiodns_init]");
            if (!aiodns_init())
            {
                Console.ReadLine();
                return;
            }
            Console.ReadLine();

            Console.WriteLine("[AIODNSTester][aiodns_free]");
            aiodns_free();
        }
    }
}
```
