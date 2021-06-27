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
  ) { }

  async create(createWordDto: CreateWordDto): Promise<{ id: string }> {
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

  async autofill(createWordDto: CreateWordDto): Promise<{ id: string }> {
    const dictionaryData = await this.lookupWord(createWordDto.word);
    return this.create({
      ...createWordDto,
      pointValue: createWordDto.pointValue ?? this.getPointValue(createWordDto.word),
      definition: createWordDto.definition ?? dictionaryData.definition,
      partOfSpeech: createWordDto.partOfSpeech ?? dictionaryData.partOfSpeech,
      synonym: createWordDto.synonym ?? dictionaryData.synonym
    });
  }

  private getPointValue(word: string): number {
    return 0;
  }

  private async lookupWord(word: string): Promise<{
    definition: string,
    partOfSpeech: string,
    synonym: string
  }> {
    return {
      definition: "",
      partOfSpeech: "",
      synonym: ""
    };
  }
}
