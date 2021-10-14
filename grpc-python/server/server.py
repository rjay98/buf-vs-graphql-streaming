
import asyncio
from datetime import datetime

from grpclib.utils import graceful_exit
from grpclib.server import Server

import json
import os
import psutil
import sys
sys.path.append('..')
from v1.seatsaver_pb2 import GetVenueResponse, GetVenuesResponse, Venue, PingResponse, PingStreamResponse
from v1.seatsaver_grpc import SeatSaverServiceBase

class SeatSaverService(SeatSaverServiceBase):

    async def GetVenues(self, stream):
        for i in range(5):
            await stream.send_message(GetVenuesResponse(venue = None, message = 'No venue found'))

    async def GetVenue(self, stream):
        request = await stream.recv_message()
        if request.venue_id == '1':
            venue = Venue(
                id = '1',
                name = 'Madison Square Garden',
                address = '4 Pennslyvania Plaza',
                city = 'New York',
                state_province = 'NY',
                postal_code = '10001',
                country = 'USA',
                changed = '2021-10-14',
                created = '2021-10-14',
                seats = []
            )
            response = GetVenueResponse(
                venue = venue,
                message = 'Found information for venue: %s' % request.venue_id
            )
            await stream.send_message(response)
        else:
            await stream.send_message(GetVenueResponse(venue = None, message = 'No venue found'))

    async def BuySeat(self, stream):
        pass

    async def GetOpenSeats(self, stream):
        pass

    async def GetReservedSeats(self, stream):
        pass

    async def GetSeats(self, stream):
        pass

    async def GetSoldSeats(self, stream):
        pass

    async def ReserveSeat(self, stream):
        pass

    async def ReleaseSeat(self, stream):
        pass

    async def Ping(self, stream):
        process = psutil.Process(os.getpid())
        runtime_info = {
            "current_time": datetime.now().isoformat(),
            "memory_usage": process.memory_info().vms
        }
        response = PingResponse(
            runtime_info = json.dumps(runtime_info),
            message = "Ping response"
        )
        await stream.send_message(response)

    async def PingStream(self, stream):
        request = await stream.recv_message()
        process = psutil.Process(os.getpid())
        item_count = request.stream_item_count or 10
        for i in range(item_count):
            runtime_info = {
                "current_time": datetime.now().isoformat(),
                "memory_usage": process.memory_info().vms
            }
            response = PingStreamResponse(
                runtime_info = json.dumps(runtime_info),
                message = f"Ping response {i}"
            )
            await stream.send_message(response)

async def main(*, host='127.0.0.1', port=50051):
    server = Server([SeatSaverService()])
    # Note: graceful_exit isn't supported in Windows
    with graceful_exit([server]):
        await server.start(host, port)
        print(f'Serving on {host}:{port}')
        await server.wait_closed()


if __name__ == '__main__':
    asyncio.run(main())
