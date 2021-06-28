import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { WordsModule } from './api/words/words.module';
import { ConfigModule, ConfigService } from '@nestjs/config';
import databaseConfig from './config/database.config';
import { Word } from './api/words/entities/word.entity';
import { PuzzlesModule } from './api/puzzles/puzzles.module';
import { RecordsModule } from './api/records/records.module';
import { RoomsModule } from './api/rooms/rooms.module';
import rapidapiConfig from './config/rapidapi.config';
import { Room } from './api/rooms/entities/room.entity';

@Module({
  controllers: [AppController],
  imports: [
    ConfigModule.forRoot({
      load: [
        databaseConfig, 
        rapidapiConfig
      ]
    }),
    TypeOrmModule.forRootAsync({
      imports: [ConfigModule],
      useFactory: (config: ConfigService) => ({
        type: 'postgres',
        host: config.get('database.host'),
        port: +config.get<number>('database.port'),
        username: config.get('database.username'),
        password: config.get('database.password'),
        database: config.get('database.database'),
        entities: [
          Word,
          Room
        ],
        synchronize: true
      }),
      inject: [ConfigService],
    }),
    WordsModule,
    PuzzlesModule,
    RecordsModule,
    RoomsModule
  ],
})
export class AppModule { }
