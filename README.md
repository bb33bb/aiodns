# aiodns
[![](https://img.shields.io/badge/Telegram-Group-blue)](https://t.me/aioCloud)
[![](https://img.shields.io/badge/Telegram-Channel-green)](https://t.me/aioCloud_channel) 

```bash
wget -O aiodns.conf https://raw.githubusercontent.com/felixonmars/dnsmasq-china-list/master/accelerated-domains.china.conf
sed -i "s/server=\///" aiodns.conf
sed -i "s/\/114.114.114.114//" aiodns.conf
```

```csharp
namespace AIODNSTester
{
    public enum NameList : int
    {
        TYPE_REST,
        TYPE_ADDR,
        TYPE_LIST,
        TYPE_CDNS,
        TYPE_ODNS
    }

    public class AIODNSTester
    {
        [DllImport("aiodns.bin", CallingConvention = CallingConvention.Cdecl)]
        public static extern bool aiodns_dial(NameList name, string value);

        [DllImport("aiodns.bin", CallingConvention = CallingConvention.Cdecl)]
        public static extern bool aiodns_init();

        [DllImport("aiodns.bin", CallingConvention = CallingConvention.Cdecl)]
        public static extern void aiodns_free();

        public static void Main(string[] args)
        {
            Console.WriteLine("[AIODNSTester][aiodns_dial]");
            aiodns_dial(NameList.TYPE_REST, "");                        // Reset
            aiodns_dial(NameList.TYPE_ADDR, ":53");                     // Listen Addr
            aiodns_dial(NameList.TYPE_LIST, "D:\\china.conf");          // China Domain Conf
            aiodns_dial(NameList.TYPE_CDNS, "tcp://119.29.29.29:53");   // China DNS
            aiodns_dial(NameList.TYPE_ODNS, "tls://1.1.1.1:853");       // Other DNS
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
