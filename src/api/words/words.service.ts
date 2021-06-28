import { ConflictException, HttpService, Injectable, InternalServerErrorException, Logger, NotFoundException } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';
import { CreateWordDto } from './dto/create-word.dto';
import { UpdateWordDto } from './dto/update-word.dto';
import { Word } from './entities/word.entity';

@Injectable()
export class WordsService {
  private readonly logger = new Logger(WordsService.name);

  constructor(
    @InjectRepository(Word)
    private wordsRepository: Repository<Word>,
    private httpService: HttpService,
    private config: ConfigService,
  ) { }

  async create(createWordDto: CreateWordDto): Promise<{ id: string }> {
    try {
      const result = await this.wordsRepository.insert(createWordDto);
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

  async remove(id: string): Promise<void> {
    const result = await this.wordsRepository.delete(id);
    if (result.affected === 0) throw new NotFoundException();
  }

  async autofill(word: string): Promise<{ id: string }> {
    word = word.toLowerCase();
    const dictionaryData = await this.lookupWord(word);
    const isPanagram = this.isPanagram(word);
    return this.create({
      word,
      pointValue: this.getPointValue(word, isPanagram),
      definition: dictionaryData.definition,
      partOfSpeech: dictionaryData.partOfSpeech,
      synonym: dictionaryData.synonym,
      isPanagram
    });
  }

  private isPanagram(word: string): boolean {
    const letterMap = new Map<string, boolean>();
    for (let char of word) {
      letterMap.set(char, true);
    }
    return [...letterMap.keys()].length >= 7;
  }

  private getPointValue(word: string, isPanagram: boolean): number {
    const wordLength = word.length;
    let score = 0;
    if (wordLength === 4) {
      score = 1;
    } else {
      score = wordLength;
    }
    if (isPanagram) score += 7;
    return score;
  }

  private async lookupWord(word: string): Promise<{
    definition?: string,
    partOfSpeech?: string,
    synonym?: string
  }> {
    const headers = {
      "x-rapidapi-key": this.config.get('rapidapi.key'),
      "x-rapidapi-host": this.config.get('rapidapi.host')
    };

    try {
      const response = await this.httpService.get(`https://wordsapiv1.p.rapidapi.com/words/${word}`, { headers }).toPromise();
  
      const result = response.data.results[0];
      
      let definition = result.definition ?? null;
      let partOfSpeech = result.partOfSpeech ?? null;
      let synonym = result.synonyms[0] ?? null;
  
      return {
        definition,
        partOfSpeech,
        synonym
      };
    } catch (err) {
      this.logger.log(err);
    }
  }
}
