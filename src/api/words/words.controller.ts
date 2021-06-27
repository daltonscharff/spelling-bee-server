import { Controller, Get, Post, Body, Patch, Param, Delete, NotFoundException, ParseUUIDPipe } from '@nestjs/common';
import { WordsService } from './words.service';
import { CreateWordDto } from './dto/create-word.dto';
import { UpdateWordDto } from './dto/update-word.dto';

@Controller('api/words')
export class WordsController {
  constructor(private readonly wordsService: WordsService) {}

  @Post()
  async create(@Body() createWordDto: CreateWordDto) {
    const result = await this.wordsService.create(createWordDto);
    return this.findOne(result.identifiers[0]["id"]);
  }

  @Get()
  findAll() {
    return this.wordsService.findAll();
  }

  @Get(':id')
  async findOne(@Param('id', new ParseUUIDPipe()) id: string) {
    const result = await this.wordsService.findOne(id);
    
    if (!result) {
      throw new NotFoundException();
    }
    return result;
  }

  @Patch(':id')
  async update(@Param('id', new ParseUUIDPipe()) id: string, @Body() updateWordDto: UpdateWordDto) {
    const result = await this.wordsService.update(id, updateWordDto);

    if (result.affected === 0) {
      throw new NotFoundException();
    }
  }

  @Delete(':id')
  async remove(@Param('id', new ParseUUIDPipe()) id: string) {
    const result = await this.wordsService.removeOne(id);
    if (result.affected === 0) {
      throw new NotFoundException();
    }
  }
}
