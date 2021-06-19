import { Injectable } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { CreateWordDto } from './dto/create-word.dto';
import { UpdateWordDto } from './dto/update-word.dto';
import { Word } from './entities/word.entity';

@Injectable()
export class WordsService {
  constructor(
    @InjectRepository(Word)
    private wordsRepository: Repository<Word>
  ) {}

  create(createWordDto: CreateWordDto) {
    // return 'This action adds a new word';
    this.wordsRepository.insert({
      word: "Test",
      pointValue: 1
    })
  }

  findAll(): Promise<Word[]> {
    return this.wordsRepository.find();
  }

  findOne(id: string) {
    return `This action returns a #${id} word`;
  }

  update(id: string, updateWordDto: UpdateWordDto) {
    return `This action updates a #${id} word`;
  }

  remove(id: string) {
    return `This action removes a #${id} word`;
  }
}
