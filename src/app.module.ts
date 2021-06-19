import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { WordsModule } from './api/words/words.module';
import { ConfigModule, ConfigService } from '@nestjs/config';
import databaseConfig from './config/database.config';
import { Word } from './api/words/entities/word.entity';

@Module({
  controllers: [AppController],
  imports: [
    ConfigModule.forRoot({
      load: [databaseConfig]
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
          Word
        ],
        synchronize: true
      }),
      inject: [ConfigService],
    }),
    WordsModule
  ],
})
export class AppModule {}
