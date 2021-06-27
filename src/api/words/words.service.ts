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
    return this.wordsRepository.insert(createWordDto);
  }

  findAll(): Promise<Word[]> {
    return this.wordsRepository.find();
  }

  findOne(id: string) {
    return this.wordsRepository.findOne(id);
  }

  update(id: string, updateWordDto: UpdateWordDto) {
    return this.wordsRepository.update(id, updateWordDto);
  }

  removeOne(id: string) {
    return this.wordsRepository.delete(id);
  }
}
