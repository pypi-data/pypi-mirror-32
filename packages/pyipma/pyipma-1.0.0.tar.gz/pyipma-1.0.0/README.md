# pyipma
Python library to retrieve information from Instituto Português do Mar e Atmosfera

## Requirements
- aiohttp
- BeautifulSoup4
- geopy

## Example

```python
import asyncio

from pyipma import *

async def main():
    async with aiohttp.ClientSession() as session:
        station = await Station.get(session, 40.6147336,-8.6424433)
        print("Nearest station if {}".format(station.local))
        print(await station.forecast())
        print(await station.observation())

if __name__ == "__main__":
    asyncio.get_event_loop().run_until_complete(main())

```
