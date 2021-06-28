import { ConflictException, Injectable, InternalServerErrorException, Logger, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { CreateRoomDto } from './dto/create-room.dto';
import { UpdateRoomDto } from './dto/update-room.dto';
import { Room } from './entities/room.entity';

@Injectable()
export class RoomsService {
  private readonly logger = new Logger(RoomsService.name);

  constructor(
    @InjectRepository(Room)
    private roomsRepository: Repository<Room>
  ) { }

  async create(createRoomDto: CreateRoomDto): Promise<{ id: string }> {
    try {
      const result = await this.roomsRepository.insert(createRoomDto);
      return {
        id: result.identifiers[0]["id"]
      };
    } catch (err) {
      switch (err.code) {
        case "23505":
          throw new ConflictException(err.detail);
        default:
          throw new InternalServerErrorException();
      }
    }
  }

  findAll(): Promise<Room[]> {
    return this.roomsRepository.find();
  }

  async findOne(id: string): Promise<Room> {
    const room = await this.roomsRepository.findOne(id);
    if (!room) throw new NotFoundException();
    return room;
  }

  async update(id: string, updateRoomDto: UpdateRoomDto): Promise<void> {
    const result = await this.roomsRepository.update(id, updateRoomDto);
    if (result.affected === 0) throw new NotFoundException();
  }

  async remove(id: string): Promise<void> {
    const result = await this.roomsRepository.delete(id);
    if (result.affected === 0) throw new NotFoundException();
  }
}
