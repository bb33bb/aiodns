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
    public enum NameList : int
    {
        TYPE_REST,
        TYPE_ADDR,
        TYPE_LIST,
        TYPE_CDNS,
        TYPE_ODNS,
        TYPE_METH
    }

    public class AIODNSTester
    {
        [DllImport("aiodns.bin", CallingConvention = CallingConvention.Cdecl)]
        public static extern bool aiodns_dial(int name, byte[] value);

        [DllImport("aiodns.bin", CallingConvention = CallingConvention.Cdecl)]
        public static extern bool aiodns_init();

        [DllImport("aiodns.bin", CallingConvention = CallingConvention.Cdecl)]
        public static extern void aiodns_free();

        public static void Main(string[] args)
        {
            Console.WriteLine("[AIODNSTester][aiodns_dial]");
            aiodns_dial((int)NameList.TYPE_REST, null);                                     // Reset
            aiodns_dial((int)NameList.TYPE_ADDR, Encoding.UTF8.GetBytes(":53"));            // Listen Addr
            aiodns_dial((int)NameList.TYPE_LIST, Encoding.UTF8.GetBytes("D:\\china.conf")); // China Domain Conf
            aiodns_dial((int)NameList.TYPE_CDNS, Encoding.UTF8.GetBytes("223.5.5.5:53"));   // China DNS
            aiodns_dial((int)NameList.TYPE_ODNS, Encoding.UTF8.GetBytes("1.1.1.1:53"));     // Other DNS
            aiodns_dial((int)NameList.TYPE_METH, Encoding.UTF8.GetBytes("TCP"));            // Method
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
