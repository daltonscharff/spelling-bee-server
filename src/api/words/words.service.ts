import { Injectable, NotFoundException } from '@nestjs/common';
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

  async create(createWordDto: CreateWordDto): Promise<{id: String}> {
    const result = await this.wordsRepository.insert(createWordDto);
    return {
      id: result.identifiers[0]["id"]
    };
  }

  findAll(): Promise<Word[]> {
    return this.wordsRepository.find();
  }

  async findOne(id: string): Promise<Word> {
    const word = await this.wordsRepository.findOne(id);
    if (!word) throw new NotFoundException();
    return word;
  }

  async update(id: string, updateWordDto: UpdateWordDto): Promise<void> {
    const result = await this.wordsRepository.update(id, updateWordDto);
    if (result.affected === 0) throw new NotFoundException();
  }

  async removeOne(id: string): Promise<void> {
    const result = await this.wordsRepository.delete(id);
    if (result.affected === 0) throw new NotFoundException();
  }

  async defineWord(id: string): Promise<void> {
    const word = this.findOne(id);
    
  }
}
