import { HttpModule, Module } from '@nestjs/common';
import { WordsService } from './words.service';
import { WordsController } from './words.controller';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Word } from './entities/word.entity';
import { ConfigModule } from '@nestjs/config';

@Module({
  imports: [
    TypeOrmModule.forFeature([Word]),
    HttpModule,
    ConfigModule
  ],
  controllers: [WordsController],
  providers: [WordsService]
})
export class WordsModule { }
